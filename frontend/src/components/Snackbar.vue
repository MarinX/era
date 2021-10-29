<template>
  <v-snackbar
      v-model="show"
  >
    {{ text }}
    <template v-slot:action="{ attrs }">
      <v-btn
          text
          v-bind="attrs"
          @click="show = false"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script>
export default {
  data: () => ({
    show: false,
    text: '',
  }),
  methods: {
    open(text) {
      this.show = true;
      this.text = text;
    },
  }
}
</script>