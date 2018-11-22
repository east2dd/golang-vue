<template>
  <div class="page-container">
    <app-header />
    <component :is="layout">
        <router-view />
    </component>
  </div>
</template>

<script>
    import {mapGetters} from 'vuex';
    import * as types from './store/types';
    import Header from './components/header/header.vue'

    const default_layout = 'default';
    
    export default {
        components: {
            'app-header': Header
        },
        computed: {
            ...mapGetters({
                counter: types.DOUBLE_COUNTER
            }),
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