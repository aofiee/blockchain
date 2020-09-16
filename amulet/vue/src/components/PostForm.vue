<template>
  <div>
    <app-text type="h2">New {{ value.type }}</app-text>
    <div v-for="field in value.fields" :key="field">
      <div v-if="field === 'Year'">
        <select v-model="myear" :disabled="flight" @change="onChangeYear($event)">
          <option disabled value>{{title(field)}}</option>
          <option v-for="(year,key) in years" :value="year" v-bind:key="key">{{ year }}</option>
        </select>
        <input type="hidden" v-model="fields[field]" />
      </div>
      <div v-else-if="field === 'AmuletType'">
        <select v-model="myear" :disabled="flight" @change="onChangeAmuletType($event)">
          <option disabled value>{{title(field)}}</option>
          <option>พระเนื้อดิน</option>
          <option>พระเนื้อชิน</option>
          <option>พระเนื้อผง</option>
          <option>พระเนื้อว่าน</option>
          <option>พระเบญจภาคี</option>
          <option>พระปิดตา</option>
          <option>พระเนื้อโลหะ ประเภทพระหล่อ พระปั้ม พระฉีด</option>
          <option>เครื่องรางของขลัง</option>
          <option>พระจากวัสดุอื่นๆ</option>
        </select>
        <input type="hidden" v-model="fields[field]" />
      </div>
      <div v-else>
        <app-input
          v-model="fields[field]"
          type="text"
          :placeholder="title(field)"
          :disabled="flight"
        />
      </div>
    </div>
    <button
      :class="['button', `button__valid__${!!valid && !flight && hasAddress}`]"
      @click="submit"
    >
      Create {{ value.type }}
      <div class="button__label" v-if="flight">
        <div class="button__label__icon">
          <icon-refresh />
        </div>Sending transaction...
      </div>
    </button>
  </div>
</template>

<style scoped>
select {
  border: none;
  font-size: inherit;
  padding: 0.75rem 2rem 0.75rem 1rem;
  margin-bottom: 0.5rem;
  width: 100%;
  max-width: 500px;
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.1);
  border-radius: 6px;
  box-sizing: border-box;
  background: rgba(0, 0, 0, 0);
  font-family: inherit;
}
button {
  background: none;
  border: none;
  color: rgba(0, 125, 255);
  padding: 0;
  font-size: inherit;
  font-weight: 800;
  font-family: inherit;
  text-transform: uppercase;
  margin-top: 0.5rem;
  cursor: pointer;
  transition: opacity 0.1s;
  letter-spacing: 0.03em;
  transition: color 0.25s;
  display: inline-flex;
  align-items: center;
}
button:focus {
  opacity: 0.85;
  outline: none;
}
.button.button__valid__true:active {
  opacity: 0.65;
}
.button__label {
  display: inline-flex;
  align-items: center;
}
.button__label__icon {
  height: 1em;
  width: 1em;
  margin: 0 0.5em 0 0.5em;
  fill: rgba(0, 0, 0, 0.25);
  animation: rotate linear 4s infinite;
}
.button.button__valid__false {
  color: rgba(0, 0, 0, 0.25);
  cursor: not-allowed;
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
      myear: "",
    };
  },
  created() {
    (this.value.fields || []).forEach((field) => {
      this.$set(this.fields, field, "");
    });
  },
  computed: {
    years() {
      const year = new Date().getFullYear();
      return Array.from(
        { length: year - 1990 },
        (value, index) => 1991 + index + 543
      );
    },
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
  methods: {
    onChangeYear(event) {
      console.log(event.target.value);
      this.fields["Year"] = event.target.value;
    },
    onChangeAmuletType(event) {
      console.log(event.target.value);
      this.fields["AmuletType"] = event.target.value;
    },
    title(string) {
      return string.charAt(0).toUpperCase() + string.slice(1);
    },
    async submit() {
      if (this.valid && !this.flight && this.hasAddress) {
        this.flight = true;
        const payload = { type: this.value.type, body: this.fields };
        await this.$store.dispatch("entitySubmit", payload);
        await this.$store.dispatch("entityFetch", payload);
        this.flight = false;
        Object.keys(this.fields).forEach((f) => {
          this.fields[f] = "";
        });
      }
    },
  },
};
</script>