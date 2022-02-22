<template>
  <div class="ion-range-slider">
    <input type="text" ref="range_input" value="" />
  </div>
</template>

<script>
import jQuery from 'jquery'
import 'ion-rangeslider'
import 'ion-rangeslider/css/ion.rangeSlider.css'
// import 'ion-rangeslider/css/ion.rangeSlider.skinHTML5.css'

export default {
  data() {
    return {
      rangeSliderObject: null
    }
  },
  model: {
    prop: 'dataRange',
    event: 'slideFinish'
  },
  props: ['dataRange'],
  mounted() {
    jQuery(this.$refs.range_input).ionRangeSlider({
      skin: 'big',
      type: 'double',
      grid: true,
      from: this.dataRange.from,
      to: this.dataRange.to,
      hide_min_max: true,
      prefix: '$',
      max_postfix: '+',
      extra_classes: '',
      onFinish: (data) => {
        // console.log('ionRangeSlider data = ', data)
        this.$store.state.priceRangeFrom = data.from
        this.$store.state.priceRangeTo = data.to
        this.$emit('slideFinish', { from: data.from, to: data.to })
      }
    })
    this.rangeSliderObject = jQuery(this.$refs.range_input).data('ionRangeSlider')
  },
  methods: {
    sayHello() {
      console.log('hi , from :', this.from)
    },
    doUpdate(min, max, from, to) {
      console.log('min=', min)
      console.log('max=', max)
      console.log('from=', from)
      console.log('to=', to)
      this.rangeSliderObject.update({
        min: min,
        max: max,
        from: from,
        to: to
      })
    }
  }

}
</script>
