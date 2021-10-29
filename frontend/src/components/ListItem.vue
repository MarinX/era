<template>
  <div>
  <v-list-item>

    <v-list-item-content>
      <v-list-item-title v-if="!selectedItem">{{item.label}}</v-list-item-title>
      <v-list-item-title  v-if="selectedItem">
        <v-form v-model="selectedItem.valid">
          <v-text-field
              label="Label"
              hide-details="auto"
              v-model="selectedItem.label"
              :rules="labelRules"
          ></v-text-field>


        </v-form>
      </v-list-item-title>
      <v-list-item-subtitle >{{item.key}}</v-list-item-subtitle>
    </v-list-item-content>


    <template v-if="selectedItem">

      <v-list-item-action>
        <v-btn icon color="green" @click="itemEdit()">
          <v-icon>mdi-check</v-icon>
        </v-btn>
      </v-list-item-action>

      <v-list-item-action>
        <v-btn icon @click="selectedItem = null;" color="red">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </v-list-item-action>


    </template>

    <template v-if="!selectedItem">
      <v-list-item-action>
        <v-btn icon @click="showQR()">
          <v-icon>mdi-cube-scan</v-icon>
        </v-btn>
      </v-list-item-action>

      <v-list-item-action>
        <v-btn icon @click="copyToClipboard()">
          <v-icon>mdi-clipboard-text</v-icon>
        </v-btn>
      </v-list-item-action>

      <v-list-item-action>
        <v-btn icon @click="selectedItem = Object.assign({}, item)">
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
      </v-list-item-action>

      <v-list-item-action>
        <v-btn icon @click="itemRemove(item)">
          <v-icon color="red">mdi-delete</v-icon>
        </v-btn>
      </v-list-item-action>
    </template>

  </v-list-item>

    <Snackbar ref="snackbar"></Snackbar>
    <QRDialog ref="qrcode"></QRDialog>
    <ConfirmDialog ref="confirm"></ConfirmDialog>
  </div>
</template>

<script>
import bridge from "@/services/bridge";
import Snackbar from '../components/Snackbar';
import QRDialog from '../components/QRDialog';
import ConfirmDialog from "@/components/ConfirmDialog";

export default  {
  components: {
    Snackbar,
    QRDialog,
    ConfirmDialog,
  },
  props: {
    item: Object,
    index: Number,
    itemKey: String,
  },
  methods:{
    copyToClipboard() {
      const item = this.$props.item;
      bridge.copyToClipboard(item.key).then(()=>{
        this.$refs.snackbar.open(`${item.label} copied to clipboard`);
      }).catch((err)=>{
        this.$refs.snackbar.open(err.message);
      })
    },
    showQR() {
      const item = this.item;
      this.$refs.qrcode.open(`Public key for ${item.label}`, item.key);
    },
    itemEdit() {
      const selectedItem = this.selectedItem;
      const index = this.$props.index;
      if(!selectedItem.valid) {
        return;
      }
      bridge["update"+this.itemKey](selectedItem.id,selectedItem.label).then(()=>{
        this.$refs.snackbar.open("Key updated");
        this.item.label = selectedItem.label;
        this.selectedItem = null;
        this.$emit('updated', {selectedItem: selectedItem, index: index});
      }).catch((err)=>{
        this.$refs.snackbar.open(err.message);
      })
    },
    itemRemove() {
      const item = this.item;
      const index = this.$props.index;
      this.$refs.confirm.open('Confirm',`Are you sure you want to remove ${item.label}?`, null).then((confirmed)=>{
        if(!confirmed) return;
        bridge["remove"+this.itemKey](item.id).then(()=>{
          this.$refs.snackbar.open(`${item.label} removed`);
          this.$emit('removed', {selectedItem: item, index: index});
        }).catch((err)=>{
          this.$refs.snackbar.open(err.message);
        });
      })
    }
  },
  data() {
    return {
      labelRules: [
        value => !!value || 'Required.',
        value => (value && value.length >= 1) || 'Min 1 character',
      ],
      selectedItem: null,
    }
  }
}
</script>