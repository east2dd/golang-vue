<template>
  <header id="header">
    <div class="logo">
      <router-link to="/">Glang Vue</router-link>
    </div>
    <nav>
      <ul>
        <li>
          <router-link to="/">Home</router-link>
        </li>
        <li>
          <router-link to="/categories">Categories</router-link>
        </li>
        <li>
          <router-link to="/products">Products</router-link>
        </li>
        <li v-if="!isAuthenticated">
          <router-link to="/signup">Sign Up</router-link>
        </li>
        <li v-if="!isAuthenticated">
          <router-link to="/signin">Sign In</router-link>
        </li>
        <li v-if="isAuthenticated">
          <router-link to="/dashboard">Dashboard</router-link>
        </li>
        <li v-if="isAuthenticated">
          <a href="#" v-on:click="signOut()">Sign Out</a>
        </li>
      </ul>
    </nav>
  </header>
</template>
<script>
    import { mapGetters } from 'vuex'
    import { AUTH_LOGOUT } from '../../store/actions/auth'

    export default {
        methods: {
            navigateToHome() {
                this.$router.push({ name: 'home' })
            },
            signOut() {
              this.$store.dispatch(AUTH_LOGOUT)
              .then(() => {
                this.$router.push({ name: 'home' })
              })
            }
        },
        computed: {
          ...mapGetters(['isAuthenticated', 'authStatus']),
          isLoggedIn() {
            return this.$store.getters.isLoggedIn;
          }
        }
    }
</script>
<style scoped>
  #header {
    height: 56px;
    display: flex;
    flex-flow: row;
    justify-content: space-between;
    align-items: center;
    background-color: #521751;
    padding: 0 20px;
  }

  .logo {
    font-weight: bold;
    color: white;
  }

  .logo a {
    text-decoration: none;
    color: white;
  }

  nav {
    height: 100%;
  }

  ul {
    list-style: none;
    margin: 0;
    padding: 0;
    height: 100%;
    display: flex;
    flex-flow: row;
    align-items: center;
  }

  li {
    margin: 0 16px;
  }

  li a {
    text-decoration: none;
    color: white!important;
  }

  li a:hover,
  li a:active,
  li a.router-link-active {
    color: #fa923f;
  }
</style>