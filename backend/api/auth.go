package api

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkl-min-1/discusspedia/backend/helper"
)

type LoginReqBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReqBody struct {
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Role      string  `json:"role" binding:"required,lowercase,oneof=siswa mahasiswa"`
	Institute string  `json:"institute" binding:"required"`
	Major     *string `json:"major" binding:"required_if=Role mahasiswa"`
	Batch     *int    `json:"batch" binding:"required_if=Role mahasiswa"`
}

type AvatarReqBody struct {
	Avatar *multipart.FileHeader `form:"avatar"`
}

type LoginSuccessResponse struct {
	Token string `json:"token"`
}

type RegisterSuccessResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var jwtKey = []byte("key")

type Claims struct {
	Id    int
	Email string
	Role  string
	jwt.StandardClaims
}

func (api *API) register(c *gin.Context) {
	var input RegisterReqBody
	err := c.BindJSON(&input)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, responseCode, err := api.userRepo.InsertNewUser(input.Name, input.Email, input.Password, input.Role, input.Institute, input.Major, input.Batch)
	if err != nil {
		c.AbortWithStatusJSON(responseCode, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(&userId, &input.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterSuccessResponse{Message: "success", Token: tokenString})
}

func (api *API) login(c *gin.Context) {
	var loginReq LoginReqBody
	err := c.BindJSON(&loginReq)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, err := api.userRepo.Login(loginReq.Email, loginReq.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role, err := api.userRepo.GetUserRole(*userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(userId, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{Token: tokenString})
}

func (api *API) changeAvatar(c *gin.Context) {
	var input AvatarReqBody
	maxFileSize := int64(1024 * 1024 * 2)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Avatar.Size > maxFileSize {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file size too large"})
		return
	}

	if !strings.Contains(input.Avatar.Header.Get("Content-Type"), "image") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "please upload an image"})
		return
	}

	userId, err := api.getUserIdFromToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userData, err := api.userRepo.GetUserData(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oldFileName := userData.Avatar

	folderPath := "media/avatar"
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	splitFilename := strings.Split(input.Avatar.Filename, ".")
	fileName := fmt.Sprintf("%s_%d.%s", userData.Name, time.Now().Unix(), splitFilename[len(splitFilename)-1])
	filePath := filepath.Join(folderPath, fileName)
	err = c.SaveUploadedFile(input.Avatar, filePath)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if oldFileName != nil {
		os.Remove(*oldFileName)
	}

	err = api.userRepo.UpdateAvatar(userId, filePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	imgUrl := fmt.Sprintf("%s://%s/%s/%s", scheme, c.Request.Host, folderPath, url.PathEscape(fileName))
	c.JSON(http.StatusOK, gin.H{"message": "success",
		"data": struct {
			Avatar string `json:"avatar"`
		}{
			Avatar: imgUrl,
		},
	})

}

func (api *API) getUserIdFromToken(c *gin.Context) (int, error) {
	tokenString := c.GetHeader("Authorization")[(len("Bearer ")):]
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return -1, err
	}

	if token.Valid {
		claim := token.Claims.(*Claims)
		return claim.Id, nil

	} else {
		return -1, errors.New("invalid token")
	}
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}

func (api API) generateJWT(userId *int, role *string) (string, error) {
	expTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Id:   *userId,
		Role: *role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
