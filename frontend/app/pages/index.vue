<script setup lang="ts"></script>
<template>
  <div class="bg-slate-50 space-y-4 md:space-y-10">
    <header
      class="items-center justify-center flex flex-col gap-6 min-h-screen"
      style="background-image: url('/landing.svg')"
    >
      <h1
        title="Fuel your body & soul - 
find recipes that taste amazing!"
        class="text-white text-5xl text-center font-bold items-center justify-center flex"
      >
        Fuel your body & soul - find recipes that taste amazing!
      </h1>
      <SearchBar />
    </header>
    <!-- some recipes -->
    <section class="space-y-6 px-6 md:px-10">
      <div class="flex justify-between">
        <h2 class="text-2xl font-semibold">Effortless Eats</h2>
        <!-- add routing  at click -->
        <button><p class="text-lg font-normal">view all recipes ></p></button>
      </div>
      <p class="text-md font-normal text-gray-600">
        Satisfy your cravings in a flash! Explore our Quick & Easy Meals for
        effortless recipes without compromising on mouthwatering taste.
      </p>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
        <RecipeCard />
        <RecipeCard />
        <RecipeCard />
        <RecipeCard />
      </div>
    </section>
    <section class="py-16 bg-white">
      <div class="container mx-auto px-4 max-w-6xl">
        <!-- Main Title -->
        <div class="text-center mb-12">
          <h1 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
            Healthy Eating Inspiration
          </h1>
          <div class="w-20 h-1 bg-green-500 mx-auto"></div>
        </div>

        <!-- Content Sections -->
        <div class="grid md:grid-cols-3 gap-8 md:gap-12">
          <!-- Section 1 -->
          <div class="text-center">
            <h2 class="text-2xl font-bold text-gray-800 mb-4">
              Delicious Dishes Packed with Nutrition
            </h2>
            <div class="w-12 h-1 bg-green-400 mx-auto mb-4"></div>
            <p class="text-gray-600">
              Explore a collection of mouthwatering recipes that not only
              delight your taste buds but also pack a powerful nutritional
              punch.
            </p>
          </div>

          <!-- Section 2 -->
          <div class="text-center">
            <h2 class="text-2xl font-bold text-gray-800 mb-4">
              Fuel Your Day with Plant-Based Power
            </h2>
            <div class="w-12 h-1 bg-green-400 mx-auto mb-4"></div>
            <p class="text-gray-600">
              Dive into the world of plant-based goodness with recipes that
              provide both energy and exceptional flavors, making each bite a
              celebration of health.
            </p>
          </div>

          <!-- Section 3 -->
          <div class="text-center">
            <h2 class="text-2xl font-bold text-gray-800 mb-4">
              One-Pan Wonders for Stress-Free Cooking
            </h2>
            <div class="w-12 h-1 bg-green-400 mx-auto mb-4"></div>
            <p class="text-gray-600">
              Simplify your culinary journey with stress-free, one-pan
              wonders—recipes that prioritize convenience without compromising
              on nutritional value or taste.
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
