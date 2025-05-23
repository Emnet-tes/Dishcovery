<script setup>
const marked = ref(false);
function toggleBookmark() {
  marked.value = !marked.value;
}
</script>
<template>
  <button
    class="absolute top-2 right-2 rounded-full"
    @click.stop.prevent="toggleBookmark"
  >
    <Icon
      name="heroicons:bookmark-20-solid"
      class="w-5 h-5 border-2 border-slate-300 rounded-full bg-white"
      v-if="marked"
    />
    <Icon
      v-else
      name="heroicons:bookmark"
      class="w-5 h-5 border border-slate-300 rounded-full bg-white"
    />
  </button>
</template>
