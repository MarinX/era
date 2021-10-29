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
        <ListItem :key="key.id" :item="key" :index="index" item-key="Contact" @updated="updateContact"  @removed="removeContact"></ListItem>
        <v-divider
            v-if="index < keys.length - 1"
            :key="index"
        ></v-divider>
      </template>
    </v-list>



  </v-container>
</template>

<script>
import bridge from "@/services/bridge";
import Snackbar from '../components/Snackbar';
import ListItem from "@/components/ListItem";

export default {
  components: {
    Snackbar,
    ListItem,
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
    createContact(data) {
      this.keys.unshift(data.item.key);
    },
    removeContact(data) {
      this.keys.splice(data.index, 1);
    },
    updateContact(data) {
      this.keys[data.index].label = data.selectedItem.label;
    },
    async loadContacts() {
      try {
        const data = await bridge.getContacts();
        console.log(data);
        if(!data.contacts) {
          this.empty = true;
        }else {
          this.keys = data.contacts;
        }
      } catch (e) {
        this.error = e.message;
      }

    }
  },
  mounted() {
    this.loadContacts();
  }
}
</script>