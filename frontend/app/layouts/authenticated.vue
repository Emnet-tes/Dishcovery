<script setup lang="ts">
import UserNavbar from '@/components/navbar/UserNavbar.vue';
</script>
<template>
    <div class="min-h-screen">
        <UserNavbar />
        <slot />
    </div>
</template>

