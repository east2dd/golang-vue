<template>
    <div>
        <h3>Product Details</h3>
        <p>ID: {{ $route.params.id }}</p>
        <router-link
                tag="button"
                :to="link"
                class="btn btn-primary">Edit Product
        </router-link>
    </div>
</template>

<script>
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
                }
            }
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