<template>
  <div class="home">
    <div class="parallax_background parallax-window"
         data-parallax="scroll"
         data-speed="0.8"></div>
    <div class="home_container">
      <div class="container">
        <div class="row">
          <div class="col">
            <div class="home_content text-center">
              <div class="home_title">
                <h1>Welcome to SwapStream</h1>
              </div>
              <div class="home_text">It's your best choice</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component({
  components: {},
})
export default class Home extends Vue {}
</script>
