<template>
  <div class="page-conVtainer">  
    <router-view name="header-top"></router-view>
    <md-app md-waterfall md-mode="fixed">
      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">
          Golang VUE
        </md-toolbar>
        <md-list>
          <md-list-item>
            <md-icon>list</md-icon>
            <span class="md-list-item-text">
              <router-link
                tag="li"
                to="/categories"
                class="list-group-item"
                style="cursor: pointer">Categories</router-link>
            </span>
          </md-list-item>
          <md-list-item>
          <md-icon>store</md-icon>
          <span class="md-list-item-text">
            <router-link
              tag="li"
              to="/products"
              class="list-group-item"
              style="cursor: pointer">Products</router-link>
          </span>
          </md-list-item>
        </md-list>
      </md-app-drawer>
      <md-app-content>
          <router-view></router-view>
      </md-app-content>      
    </md-app>
  </div>
</template>

<script>
  export default {
  }
</script>

<style lang="scss" scoped>
  .md-app {
    border: 1px solid rgba(#000, .12);
  }

   // Demo purposes only
  .md-drawer {
    width: 230px;
    max-width: calc(100vw - 125px);
  }
</style>