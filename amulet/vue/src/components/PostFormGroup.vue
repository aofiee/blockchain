<template>
  <div>
    <router-link to="/">List All Amulets</router-link>
    <div class="type" v-for="type in typeList" :key="type.type">
      <post-form :value="type" />
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    typeList() {
      return this.$store.state.app.types || [];
    },
  },
};
</script>
