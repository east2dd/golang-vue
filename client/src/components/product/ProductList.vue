<template>
  <div>
    <router-link
            tag="li"
            to="/products/new"
            class="button"
            style="cursor: pointer">New Product</router-link>
        
    <md-list class="md-double-line md-dense">

        <md-list-item v-for="item in items" :key="item.ID">
            <div class="md-list-item-text">
            <span>
                <router-link
                    tag="li"
                    :to="'/products/' + item.ID"
                    class="list-group-item"
                    style="cursor: pointer"><a href="javascript:void(0)">{{ item.Name }}</a></router-link>
            </span>
            <span>id: {{ item.ID }}</span>
            <span>description: {{ item.Description }}</span>
            </div>
        </md-list-item>
    </md-list>
  </div>
</template>
<script>
import axios from '../../axios-auth';
export default {
    data() {
        return {
            id: this.$route.params.id,
            items: null
        }
    },
    methods: {
        getProducts () {
            axios.get('/api/products')
            .then((res) => {
                this.items = res.data.data;
            })
            .catch(error => window.console.log(error))
        }
    },
    mounted () {
        this.getProducts()
    }
}
</script>