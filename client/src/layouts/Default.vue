<template>
  <div class="page-container">
    <router-view></router-view> 
  </div>
</template>

<script>
    import {mapGetters} from 'vuex';
    import Header from '../components/Header';
    import * as types from '../store/types';
    export default {
        components:{
          Header,
        },
        computed: {
        }
    }
</script>