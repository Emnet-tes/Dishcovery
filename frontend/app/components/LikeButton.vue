<script setup lang="ts">
const liked = ref(false);

const toggleLike = () => {
  liked.value = !liked.value;
};
</script>

<template>
   <!-- Like Button -->
   <button
          @click.stop.prevent="toggleLike"
          class="absolute bottom-2 right-2    "
        >
        <Icon
          v-if="liked"
          name="heroicons:heart-20-solid"
          class="w-5 h-5 bg-red-500"
        />
        <Icon
          v-else
          name="heroicons:heart"
          class="w-5 h-5"
        />
        </button>
</template>


