<template>
    <div>
        <h5>Product Details</h5>
        <p>ID: {{ $route.params.id }}, Name: {{ item.Name }} </p>
        <p>Price: {{ item.Price }}</p>
        <p>Description: {{ item.Description }}</p>
        <router-link
                tag="button"
                :to="link"
                class="success button">Edit Product
        </router-link>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import ProductItem from '../product/ProductItem'
  export default {
      data() {
          return {
            link: {
                name: 'productEdit',
                params: {
                    id: this.$route.params.id
                },
                hash: ''
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