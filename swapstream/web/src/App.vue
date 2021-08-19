<template>
  <div class="super_container">
    <router-view />
    <footer class="footer">
      <div class="container">
        <div class="row">
          <div class="col">
            <div class="footer_content d-flex flex-md-row flex-column align-items-center align-items-start justify-content-start">
              <div class="copyright">
              </div>
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>

</template>

<script>
export default {
  name: 'App',
  methods: {
    getName () {
      console.log('name=' + this.$store.state.name)
    }
  }
}
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
