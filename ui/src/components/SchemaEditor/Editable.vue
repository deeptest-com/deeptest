<template>
  <span
      ref="editable"
      contenteditable
      placeholder="name"
      v-on="listeners"
  />
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: 'name',
    },
  },
  computed: {
    listeners() {
      return { ...this.$listeners, input: this.onInput };
    },
  },
  mounted() {
    this.$refs.editable.innerText = this.value;
  },
  methods: {
    onInput(e) {
      this.$emit('input', e.target.innerText);
    },
  },
};
</script>
<style lang="less" scoped>
span{
  border:none;
  outline:none;
  display: inline-block;
  min-width: 20px;
}
//span[placeholder]:empty:focus::before{
//  content:none;
//}
//span:empty:before{
//  content:attr(placeholder);
//  //color:gray;
//  //margin-left: 5%;
//}


</style>
