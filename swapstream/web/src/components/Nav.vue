<template>
  <nav class="main_nav">
    <ul class="d-flex flex-row align-items-center justify-content-start">
      <li><a href="blog.html">Docs</a></li>
      <li><a href="contact.html">Vote</a></li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'Navigate',
  props: {
    msg: String
  }
}
</script>
