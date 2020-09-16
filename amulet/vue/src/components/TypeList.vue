<template>
  <div>
    <router-link to="/post">New Amulet</router-link>
    <div class="type" v-for="type in typeList" :key="type.type">
      <type-item :value="type" />
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
