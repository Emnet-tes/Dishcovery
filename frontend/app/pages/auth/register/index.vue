<script setup lang="ts">
import RegisterForm from "@/components/RegisterForm.vue";

</script>
<template>
  <div
    class="min-h-screen flex bg-cover bg-center"
    style="background-image: url('/landing.svg')"
  >
    <!-- Left Side - Image with Overlay Text -->
    <div class="hidden lg:block relative w-1/2">
      <div
        class="absolute inset-0 bg-black bg-opacity-30 flex items-center justify-center p-12"
      >
        <div class="text-white max-w-md">
          <h1 class="text-4xl font-bold mb-4">Kitchen Tales</h1>
          <p class="text-xl">
            Embark on a culinary journey with us!<br />
            Sign up to unlock a world of delicious recipes, and personalized
            cooking experiences.
          </p>
        </div>
      </div>
    </div>
    <RegisterForm />
  </div>
</template>
