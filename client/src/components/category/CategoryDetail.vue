<template>
    <div>
        <h3>Some Category Details</h3>
        <p>Category loaded has ID: {{ $route.params.id }}</p>
        <router-link
                tag="button"
                :to="link"
                class="btn btn-primary">Edit Category
        </router-link>
    </div>
</template>

<script>
  import axios from '../../axios-auth';
  import * as types from '../../store/types';
    export default {
        data() {
            return {
                link: {
                    name: 'categoryEdit',
                    params: {
                        id: this.$route.params.id
                    },
                    query: {
                        q: 100
                    },
                    hash: '#data'
                },
                id: this.$route.params.id,
                categories: null
            }
        },
        methods: {
            getCategories () {
                axios.get('/api/categories')
                .then((res) => {

                    this.$router.push({ name: 'categoryList' });
                })
                .catch(error => console.log(error))
            }
        },
         mounted () {
            this.getCategories()
        },
        beforeRouteEnter(to, from, next) {
            if (true) {
                next();
            } else {
                next(false);
            }
        }
    }
</script>