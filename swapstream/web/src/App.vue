<template>
  <div class="super_container">
    <!-- Header -->
    <header class="header">
      <div class="container">
        <div class="row">
          <div class="col">
            <div
              class="
                header_content
                d-flex
                flex-row
                align-items-center
                justify-content-start
              "
            >
              <div class="logo">
                <a href="#">
                  <div>SwapStream</div>
                  <div></div>
                </a>
              </div>
              <nav class="main_nav">
                <ul
                  class="
                    d-flex
                    flex-row
                    align-items-center
                    justify-content-start
                  "
                >
                  <li><a href="blog.html">Docs</a></li>
                  <li><a href="contact.html">Vote</a></li>
                </ul>
              </nav>
              <div
                class="
                  header_extra
                  d-flex
                  flex-row
                  align-items-center
                  justify-content-start
                  ml-auto
                "
              >
                <div class="book_button trans_200">
                  <a href="#">Launch App</a>
                </div>
              </div>
              <div class="hamburger ml-auto">
                <i class="fa fa-bars" aria-hidden="true"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Menu -->

    <div class="menu">
      <div
        class="background_image"
        style="background-image: url(images/menu.jpg)"
      ></div>
      <div
        class="
          menu_content
          d-flex
          flex-column
          align-items-center
          justify-content-center
        "
      >
        <ul class="menu_nav_list text-center">
          <li><a href="index.html">Docs</a></li>
          <li><a href="about.html">Vote</a></li>
        </ul>
        <div class="menu_review"><a href="#">Launch App</a></div>
      </div>
    </div>

    <!-- Home -->

    <div class="home">
      <div
        class="parallax_background parallax-window"
        data-parallax="scroll"
        data-speed="0.8"
      ></div>
      <div class="home_container">
        <div class="container">
          <div class="row">
            <div class="col">
              <div class="home_content text-center">
                <div class="home_title"><h1>Welcome to SwapStream</h1></div>
                <div class="home_text">It's your best choice</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
//import HelloWorld from "./components/Home.vue";

export default {
  name: "App",
  components: {},
};
</script>

<style>
/* @import "./assets/styles/bootstrap-4.1.2/bootstrap.min.css";
@import "./assets/plugins/font-awesome-4.7.0/css/font-awesome.min.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/owl.carousel.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/owl.theme.default.css";
@import "./assets/plugins/OwlCarousel2-2.2.1/animate.css";
@import "./assets/plugins/colorbox/colorbox.css";
@import "./assets/styles/main_styles.css";
@import "./assets/styles/responsive.css"; */
</style>
