<template>
  <div>
    <v-btn
        elevation="2"
        fab
        color="primary"
        fixed
        bottom
        right
        large
        @click="show = true"
    >
      <v-icon>mdi-plus</v-icon>
    </v-btn>
  <v-dialog
      v-model="show"
      persistent
      max-width="600px"
  >
    <v-card>
      <v-card-title>
        <span class="text-h5">Create Key</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="Label*"
                  ref="createKeyLabel"
                  v-model="input"
                  :rules="labelRules"
                  required
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>

      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            text
            @click="show = false"
        >
          Close
        </v-btn>
        <v-btn
            color="blue darken-1"
            text
            @click="createKey()"
        >
          Save
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
    <Snackbar ref="snackbar"></Snackbar>
  </div>
</template>
<script>
import bridge from "@/services/bridge";
import Snackbar from "@/components/Snackbar";

export default {
  components: {
    Snackbar,
  },
  props: {
    defaultValue: String,
  },
  methods: {
    createKey() {
      const key = this.$refs.createKeyLabel;
      if(!key.valid) {
        return;
      }
      const label = this.input;
      bridge.createKey(label).then((data)=>{
        this.$refs.snackbar.open(`${label} created`);
        this.show = false;
        this.$emit('created', {item: data});
      }).catch((err)=>{
        this.$refs.snackbar.open(err.message);
      });
    },
  },
  data() {
    return {
      input: this.$props.defaultValue,
      show: false,
      labelRules: [
        value => !!value || 'Required.',
        value => (value && value.length >= 1) || 'Min 1 character',
      ],
    }
  }
}
</script>