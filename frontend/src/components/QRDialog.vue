<template>
  <v-dialog
      v-model="show"
      width="500"
  >
    <v-card>
      <v-card-title class="text-center">
        {{text}}
      </v-card-title>

      <qrcode-vue :value="value" :size="300" level="H" class="text-center" />

      <v-divider></v-divider>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="primary"
            text
            @click="show = false"
        >
          Close
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import QrcodeVue from 'qrcode.vue'
export default {
  components: {
    QrcodeVue
  },
  data: () => ({
    show: false,
    text: '',
    value: '',
  }),
  methods: {
    open(text, value) {
      this.show = true;
      this.text = text;
      this.value = value;
    },
  }
}
</script>