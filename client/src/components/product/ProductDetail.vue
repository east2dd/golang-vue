<template>
    <div>
        <h5>Product Details</h5>
        <p>ID: {{ $route.params.id }}, Name: {{ item.Name }} </p>
        <router-link
                tag="button"
                :to="link"
                class="btn btn-primary">Edit Product
        </router-link>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import * as types from '../../store/types';
  import ProductItem from '../product/ProductItem'
  export default {
      data() {
          return {
            link: {
                name: 'productEdit',
                params: {
                    id: this.$route.params.id
                },
                query: {
                    locale: 'en',
                    q: 100
                },
                hash: '#data'
            },
            id: this.$route.params.id,
            item: {},
          }
      },
      methods: {
          getProduct () {
              axios.get('/api/products/' + this.id)
              .then((res) => {
                  this.item = res.data.data;
              })
              .catch(error => console.log(error))
          },
      },
      mounted () {
          this.getProduct()
      }
  }
</script>