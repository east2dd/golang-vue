<template>
    <div>
        <h3>Edit the Category</h3>
        <p>Name: {{ item.Name }}</p>
        <p>Description: {{ item.Description }}</p>
    </div>
</template>

<script>
    export default {
        data() {
          return {
              confirmed: false
          }
        },
        beforeRouteLeave(to, from, next) {
            if (this.confirmed) {
                next();
            } else {
                if (confirm('Are you sure?')) {
                    next();
                } else {
                    next(false);
                }
            }
        }
    }
</script>