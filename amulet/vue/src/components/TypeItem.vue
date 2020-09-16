<template>
  <div>
    <app-text type="h2">List of {{ value.type }} items</app-text>
    <div class="item" v-for="instance in instanceList" :key="instance.id">
      <div class="item__field" v-for="(value, key) in instance" :key="key">
        <div class="item__field__key">{{ key }}:</div>
        <div class="item__field__value">{{ value }}</div>
      </div>
    </div>
    <div
      class="card__empty"
      v-if="instanceList.length < 1"
    >There are no {{ value.type }} items yet. Create one using the form below.</div>
  </div>
</template>

<style scoped>
.item {
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
  padding: 1rem;
  border-radius: 0.5rem;
  overflow: hidden;
}
.item__field {
  display: grid;
  line-height: 1.5;
  grid-template-columns: 15% 1fr;
  grid-template-rows: 1fr;
  word-break: break-all;
}
.item__field__key {
  color: rgba(0, 0, 0, 0.25);
  word-break: keep-all;
  overflow: hidden;
}
.card__empty {
  margin-bottom: 1rem;
  border: 1px dashed rgba(0, 0, 0, 0.1);
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box;
  border-radius: 8px;
  color: rgba(0, 0, 0, 0.25);
  text-align: center;
  min-height: 8rem;
}
@keyframes rotate {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(-360deg);
  }
}
@media screen and (max-width: 980px) {
  .narrow {
    padding: 0;
  }
}
</style>

<script>
export default {
  props: ["value"],
  data: function () {
    return {
      fields: {},
      flight: false,
    };
  },
  computed: {
    hasAddress() {
      return !!this.$store.state.account.address;
    },
    instanceList() {
      return this.$store.state.data[this.value.type] || [];
    },
    valid() {
      return Object.values(this.fields).every((el) => {
        return el.trim().length > 0;
      });
    },
  },
};
</script>
