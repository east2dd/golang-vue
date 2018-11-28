<template>
    <div class="product">
        <div class="row">
            <div class="columns medium-4">
                <router-link
                    tag="p"
                    :to="'/products/' + this.item.ID"
                    style="cursor: pointer"><a href="javascript:void(0)">{{ this.item.Name }}</a></router-link>
            </div>

            <div class="columns medium-8">
                <p>Price: {{ this.item.Price }}</p>
                <p>{{ this.item.Description }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: ['item'],
}
</script>

<style scoped>
  .product {
      margin-bottom: 2em;
  }
</style>