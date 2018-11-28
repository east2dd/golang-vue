<template>
<div>
    <router-link
            tag="li"
            to="/products/new"
            class="button"
            style="cursor: pointer">New Product</router-link>
        
    <div class="row">
        <div class="columns medium-4" v-for="item in items" :key="item.ID">
            <p><img src="https://via.placeholder.com/300.png" style="width: 300px;"/></p>
            <router-link
                tag="p"
                :to="'/products/' + item.ID"
                class="list-group-item"
                style="cursor: pointer"><a href="javascript:void(0)">{{ item.Name }}</a></router-link>
            <p>{{ item.Description }}</p>
            <p>price: {{ item.Price }}</p>
        </div>
    </div>
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