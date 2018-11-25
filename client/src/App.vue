<template>
  <div class="page-container">
    <app-header />
    <component :is="layout">
        <router-view />
    </component>
  </div>
</template>

<script>
    import axios from 'axios'
    import Header from './components/header/header.vue'
    import AUTH_LOGOUT from './store/actions/auth'

    const default_layout = 'default';
    
    export default {
        components: {
            'app-header': Header
        },
        created: function () {
          axios.interceptors.response.use(undefined, function (err) {
            return new Promise(function () {
              if (err.status === 401 && err.config && !err.config.__isRetryRequest) {
                this.$store.dispatch(AUTH_LOGOUT)
              }
              throw err;
            });
          });
        },
        computed: {
            layout(){
                let layout = "";
                
                if(this.$route.matched.some(m => m.meta.layout))
                {
                    let matched = this.$route.matched.find(m => m.meta.layout)
                    layout = matched.meta.layout;
                }

                return (layout || default_layout) + '-layout';
            }
        }
    }
</script>