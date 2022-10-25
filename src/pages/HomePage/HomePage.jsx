import React from "react";
import "../../App.scss";
import "./HomePage.scss";
import Banner from "../../images/banner.svg";
import Basis from "../../images/basis.svg";
import Logo from "../../images/logo_app.png";
import Button from "../../components/Button/Button";
import Post from "../../images/post.svg";
// import Survey from "../../images/survey.svg";
// import Event from "../../images/event.svg";
import { Link } from "react-router-dom";

const HomePage = () => {
  return (
    <div className="home-page">
      <div className="container">
        <div className="words">
          <div className="title">
            <h1>Tingkatkan kehidupan sekolah Anda</h1>
            <p>Bergabunglah dengan siswa lain di seluruh Indonesia. Apakah Anda ingin menanyakan sesuatu tentang sekolah? Cukup posting dan pengguna lain akan menjawabnya.</p>
            <div className="btn">
              <Link to="/forum">
                <Button variant="login">Buka Forum</Button>
              </Link>
            </div>
          </div>
        </div>
        <div className="banner">
          <img src={Banner} alt="" />
        </div>
      </div>
      <div className="what">
        <div className="gambar">
          <img src={Basis} alt="" />
        </div>
        <div className="explain">
          <h2>Buka forum posting atau tempatkan tautan survei Anda </h2>
          <p>Kami mencoba membuat platform forum yang dapat digunakan oleh siswa di seluruh Indonesia. Mereka dapat bertanya tentang apa saja, berbagi tentang apa saja, dan memperluas koneksi mereka dengan siswa lain.</p>
          <br />
          <p>Kami juga menyediakan tempat yang bisa digunakan untuk berbagi link survey. Dengan menempatkan tautan survei Anda di BASIS, pengguna lain dapat mengisi tautan survei Anda dan Anda tidak perlu mencari responden lain.</p>
          <br />
          <p>kami berencana untuk memperluas fitur kami sehingga kami dapat menyediakan banyak layanan yang dapat membuat kehidupan sekolah atau kampus Anda lebih mudah. Jadi tetap disini.</p>
        </div>
      </div>
      <div className="fitur">
        <div className="text-fitur">
          <h1>Layanan kami</h1>
          <p>Kami memudahkan pengguna untuk menggunakan platform kami</p>
        </div>
        <div className="fitur-wrapper">
          <div className="card">
            <Link to="/forum">
              <img src={Post} alt="" />
              <h3>Post</h3>
            </Link>
          </div>
          
          {/* <div className="card">
            <img src={Survey} alt="" />
            <h3>Survey</h3>
          </div> */}
        </div>
        <div className="footer">
          <p>copyright 2022 | Tim PKL</p>
        </div>
      </div>
    </div>
  );
};

export default HomePage;
