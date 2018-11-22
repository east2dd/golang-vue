<template>
    <div>
        <h5>ID: {{ $route.params.id }}, Name: {{ item.name }}</h5>

        <md-subheader>PRODUCTS</md-subheader>

        <md-list class="md-double-line md-dense">
            <product-item v-for="item in products" :item="item" :key="item.ID" ></product-item>
        </md-list>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import * as types from '../../store/types';
  import ProductItem from '../product/ProductItem'

    export default {
        data() {
            return {
                id: this.$route.params.id,
                item: {},
                products: []
            }
        },
        components: {
            'product-item': ProductItem
        },
        methods: {
            getCategory () {
                axios.get('/api/categories/' + this.id)
                .then((res) => {
                    this.item = res.data.data;
                })
                .catch(error => console.log(error))
            },
            getProducts(){
               axios.get('/api/categories/' + this.id + '/products')
                .then((res) => {
                    this.products = res.data.data;
                })
                .catch(error => console.log(error)) 
            }
        },
        mounted () {
            this.getCategory()
            this.getProducts()
        }
    }
</script>