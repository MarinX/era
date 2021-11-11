<template>
  <v-app>
    <v-app-bar app flat>
      <v-toolbar-title></v-toolbar-title>

      <v-tabs centered class="ml-n9">
        <v-tab to="/contacts">
          Contacts
        </v-tab>
        <v-tab to="/keys">
          Keys
        </v-tab>
        <v-tab to="/wizard/encrypt">
          Encrypt
        </v-tab>
        <v-tab to="/wizard/decrypt">
          Decrypt
        </v-tab>
      </v-tabs>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-settings</v-icon>
      </v-btn>

    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>

  </v-app>
</template>



<style>
  .logo {
    width: 16em;
  }
</style>
