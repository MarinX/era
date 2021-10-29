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


    <Snackbar ref="snackbar"></Snackbar>

    <v-list v-if="error == null && !empty" elevation="5" two-line>
      <template v-for="(key,index) in keys">
      <ListItem :key="key.id" :item="key" :index="index" item-key="Key" @updated="updateKey"  @removed="removeKey"></ListItem>
        <v-divider
            v-if="index < keys.length - 1"
            :key="index"
        ></v-divider>
      </template>
    </v-list>

    <CreateKey default-value="My Key" :open="createKeyShow" @created="createKey"></CreateKey>

  </v-container>
</template>

<script>
import bridge from "@/services/bridge";
import Snackbar from '../components/Snackbar';
import ListItem from "@/components/ListItem";
import CreateKey from "@/components/CreateKey";

export default {
  components: {
    Snackbar,
    ListItem,
    CreateKey,
  },
  data() {
    return {
      labelRules: [
      value => !!value || 'Required.',
      value => (value && value.length >= 1) || 'Min 1 character',
    ],
      keys: [],
      error: null,
      empty: false,
      createKeyShow: false
    }
  },
  methods: {
    createKey(data) {
      this.keys.unshift(data.item.key);
    },
    removeKey(data) {
      this.keys.splice(data.index, 1);
    },
    updateKey(data) {
      this.keys[data.index].label = data.selectedItem.label;
    },
    async loadKeys() {
      try {
        const data = await bridge.getKeys();
        console.log(data);
        if(!data.keys) {
          this.empty = true;
        }else {
          this.keys = data.keys;
        }
      } catch (e) {
        this.error = e.message;
      }

    }
  },
  mounted() {
    this.loadKeys();
  }
}
</script>