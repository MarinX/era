<template>
  <v-container>

    <div v-if="error">
      <v-alert
          color="red lighten-2"
          class="text-center"
      >
        {{error}}
      </v-alert>
    </div>


    <v-stepper v-model="e1">
      <v-stepper-header>
        <v-stepper-step
            :complete="e1 > 1"
            step="1"
        >
          Select type
        </v-stepper-step>

        <v-divider></v-divider>

        <v-stepper-step
            :complete="e1 > 2"
            step="2"
        >
          Select key
        </v-stepper-step>

        <v-divider></v-divider>

        <v-stepper-step step="3">
          Results
        </v-stepper-step>
      </v-stepper-header>

      <v-stepper-items>
        <v-stepper-content step="1">
          <v-card class="mb-5">
            <v-card-title>What do you want to decrypt?</v-card-title>
            <v-card-text>
              <v-radio-group mandatory v-model="selectedType" v-on:change="selectedInput=''">
                <v-radio
                    label="File"
                    value="file"
                ></v-radio>

                <v-text-field
                    v-show="selectedType === 'file'"
                    show-size
                    truncate-length="50"
                    label="Select file"
                    v-on:click="selectFile"
                    v-model="selectedInput"
                ></v-text-field>

                <v-radio
                    label="Text"
                    value="text"
                ></v-radio>

                <v-textarea
                    v-show="selectedType === 'text'"
                    solo
                    v-model="selectedInput"
                    label="Enter text to decrypt"
                    row-height="75"
                    auto-grow
                ></v-textarea>
              </v-radio-group>

            </v-card-text>
          </v-card>
          <v-btn
              color="primary"
              @click="e1 = 2"
              :disabled="selectedInput.length === 0"
          >
            Next
          </v-btn>
        </v-stepper-content>

        <v-stepper-content step="2">
          <v-card
              class="mb-5"
          >
            <v-card-title>Select key to decrypt</v-card-title>
            <v-card-text>
              <v-combobox
                  :item-text="(obj) => {return obj.label}"
                  v-model="selectedKeys"
                  :items="keys"
                  label="Select key/s"
                  multiple
              ></v-combobox>
            </v-card-text>
          </v-card>

          <v-btn
              color="primary"
              :disabled="selectedKeys.length === 0"
              @click="e1 = 3; processEncrypt()"
          >
            Next
          </v-btn>

          <v-btn text @click="e1=1">
            Back
          </v-btn>
        </v-stepper-content>

        <v-stepper-content step="3">
          <v-card
              class="mb-5"
          >
            <v-card-title>Your results</v-card-title>
            <v-card-text>
              <v-textarea
                  solo
                  v-model="resultText"
                  readonly
              ></v-textarea>
            </v-card-text>
          </v-card>


        </v-stepper-content>
      </v-stepper-items>
    </v-stepper>



  </v-container>
</template>
<script>
import bridge from "@/services/bridge";

export default {

  data() {
    return {
      e1: 1,
      selectedInput: '',
      resultText: '',
      selectedType: 'file',
      selectedKeys: [],
      keys: [],
      keysEmpty: false,
      error: null,
      loading: true,
    }
  },
  methods: {
    async selectFile() {
      try {
        const data = await bridge.selectFile();
        console.log(data);
        this.selectedInput = data;

      } catch (e) {
        this.error = e.message;
      }
    },
    async loadKeys() {
      try {
        const data = await bridge.getKeys();
        if(!data.keys) {
          this.keysEmpty = true;
        }else {
          this.keys = data.keys;
          console.log(this.keys);
        }
      } catch (e) {
        this.error = e.message;
      }
    },

    async processEncrypt() {
      try{
        this.resultText = await bridge.decrypt(this.selectedType, this.selectedKeys.map((c)=>{return c.id}),this.selectedInput);
      }catch (e) {
        this.error = e.message;
      }

    }
  },
  mounted() {
    this.loadKeys();
  }
}

</script>