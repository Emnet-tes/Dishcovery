<template>
    <div class="container mx-auto px-4 py-8">
        <div class="max-w-4xl mx-auto">
            <!-- Breadcrumbs -->
            <nav class="text-sm text-gray-500 mb-4">
                <ol class="list-none p-0 inline-flex">
                    <li class="flex items-center">
                        <NuxtLink to="/" class="text-gray-500 hover:text-gray-700">Home</NuxtLink>
                        <svg class="fill-current w-3 h-3 mx-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512">
                            <path
                                d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 67.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z" />
                        </svg>
                    </li>
                    <li class="flex items-center">
                        <NuxtLink to="/recipes" class="text-gray-500 hover:text-gray-700">Search</NuxtLink>
                        <svg class="fill-current w-3 h-3 mx-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512">
                            <path
                                d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 67.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z" />
                        </svg>
                    </li>
                    <li class="flex items-center text-gray-700">
                        Kung Pao Chicken
                    </li>
                </ol>
            </nav>

            <!-- Recipe Header -->
            <h1 class="text-3xl font-bold text-gray-800 mb-2">Kung Pao Chicken</h1>
            <div class="flex items-center text-sm text-gray-600 mb-4">
                By <span class="ml-1 font-semibold">Judy Leung</span> <span class="mx-2">|</span> January 10, 2024
            </div>
            <div class="flex items-center text-sm text-gray-600 mb-4">
                <div class="flex items-center mr-4">
                    <svg class="w-4 h-4 text-yellow-500 mr-1" fill="currentColor" viewBox="0 0 20 20">
                        <path
                            d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.83-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z">
                        </path>
                    </svg>
                    <span class="font-semibold">4.6</span> (27)
                </div>
            </div>

            <!-- Recipe Image -->
            <img src="/landing.svg" alt="Kung Pao Chicken"
                class="w-full h-96 object-cover rounded-lg mb-6">

            <!-- Action Buttons -->
            <div class="flex items-center justify-center space-x-4 mb-8">
                <button
                    class="flex items-center px-4 py-2 border rounded-full text-gray-700 hover:bg-gray-100 transition-colors">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                    </svg>
                    SAVE
                </button>
                <button
                    class="flex items-center px-4 py-2 border rounded-full text-gray-700 hover:bg-gray-100 transition-colors">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.539 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.83-.197-1.539-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                    </svg>
                    RATE
                </button>
                <button
                    class="flex items-center px-4 py-2 border rounded-full text-gray-700 hover:bg-gray-100 transition-colors">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6m4 0H9m7 0l3-3m0 0l-3-3m0 3H9" />
                    </svg>
                    PRINT
                </button>

                <!-- Social Share Icons (Placeholders) -->
                <div class="flex items-center space-x-2 ml-4">
                    <a href="#" class="text-gray-500 hover:text-gray-700"><i class="fab fa-facebook"></i></a>
                    <a href="#" class="text-gray-500 hover:text-gray-700"><i class="fab fa-twitter"></i></a>
                    <a href="#" class="text-gray-500 hover:text-gray-700"><i class="fab fa-pinterest"></i></a>
                    <a href="#" class="text-gray-500 hover:text-gray-700"><i class="fas fa-envelope"></i></a>
                </div>
            </div>

            <!-- Overview -->
            <div class="mb-8">
                <h2 class="text-xl font-bold text-gray-800 mb-2">Overview</h2>
                <p class="text-gray-700 leading-relaxed">
                    Kung Pao Chicken is a Chinese takeout favourite that is mouthwateringly good and highly addictive –
                    so it's a good thing it's easy to make at home so we don't need to order takeout every time we crave
                    it! We love the strong flavoured sweet-sour-savoury sauce with the signature tingle of numbing heat
                    from the Sichuan pepper!
                </p>
                <p class="text-gray-700 leading-relaxed mt-4">
                    If you're wondering whether Kung Pao Chicken is authentic Chinese, the dish as we know it outside of
                    China is a slightly westernised version of an authentic Chinese Sichuan dish. Traditional Kung Pao
                    Chicken is a dry stir fry, which means, unlike 99% of other Asian stir fries, it's not swimming in
                    loads of sauce. But with Kung Pao Chicken, the sauce is very intense flavoured so you don't need
                    loads of it. When it mixes in with the rice, just a bit of sauce goes a long way.
                </p>
                <p class="text-gray-700 leading-relaxed mt-4">
                    There are good reasons why everyone loves kung pao chicken. It's got so many flavors going on:
                    tangy, sweet, and salty with a hint of heat. The art is putting the right amount of each ingredient
                    to come up with that winning flavor combination.
                </p>
                <p class="text-gray-700 leading-relaxed mt-4">
                    It's actually a relatively easy dish to make at home, and my goal was to make a restaurant-quality
                    kung pao chicken recipe. Try it for yourself to see if I've succeeded!
                </p>
            </div>

            <!-- Cook's Tips -->
            <div class="mb-8">
                <h2 class="text-xl font-bold text-gray-800 mb-2">Cook's Tips</h2>
                <ul class="list-disc list-inside text-gray-700 leading-relaxed space-y-2">
                    <li>As with all stir-fries, the dish cooks quickly, so be sure to have all your ingredients prepped
                        before you start cooking.</li>
                    <li>Cook the Kung Pao sauce down until it reduces to a syrupy consistency with quite an intense
                        flavour. That's the Kung Pao way!!</li>
                    <li>Chinese Cooking Wine, also known as Shaoxing wine or Shaosing wine, is the secret ingredient
                        that makes homemade Chinese food truly taste as good as take out.</li>
                </ul>
            </div>

            <!-- Timing and Servings -->
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-center mb-8">
                <div class="p-4 border rounded-lg">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mx-auto text-gray-500 mb-2" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <p class="text-sm text-gray-600">Prep Time</p>
                    <p class="font-semibold">30 mins</p>
                </div>
                <div class="p-4 border rounded-lg">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mx-auto text-gray-500 mb-2" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <p class="text-sm text-gray-600">Cook Time</p>
                    <p class="font-semibold">10 mins</p>
                </div>
                <div class="p-4 border rounded-lg">
                    <svg xmlns="http://www.w3.org/3516/svg" class="h-6 w-6 mx-auto text-gray-500 mb-2" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <p class="text-sm text-gray-600">Total Time</p>
                    <p class="font-semibold">40 mins</p>
                </div>
                <div class="p-4 border rounded-lg">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mx-auto text-gray-500 mb-2" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M18 18.72a9.935 9.935 0 003.975-.99A11.25 11.25 0 0012 21.75c-4.505 0-8.422-2.08-11.177-5.34M12 12v7.5m6-6h3m-3 0a9.935 9.935 0 00-3.975-.99M10.5 6H6m4.5 0A4.5 4.5 0 1115 10.5m-4.5-4.5v4.5m0-4.5H6m11.25 4.5a4.5 4.5 0 100 9M12.75 12.75h6M12.75 12.75V18" />
                    </svg>
                    <p class="text-sm text-gray-600">Servings</p>
                    <p class="font-semibold">4</p>
                </div>
            </div>

            <!-- Ingredients -->
            <div class="mb-8">
                <div class="flex justify-between items-center mb-4">
                    <h2 class="text-xl font-bold text-gray-800">Ingredients</h2>
                    <div>
                        <button class="text-sm font-medium text-gray-600 hover:text-gray-800">METRIC</button>
                        <span class="text-gray-400">|</span>
                        <button class="text-sm font-medium text-gray-800">US</button>
                    </div>
                </div>

                <!-- Ingredient List (Placeholder) -->
                <div class="space-y-2 text-gray-700 leading-relaxed">
                    <p><strong>To Roast the Peanuts:</strong></p>
                    <ul class="list-disc list-inside ml-4 space-y-1">
                        <li>1 teaspoon vegetable oil</li>
                        <li>180 g raw peanuts (shelled, with or without the skin)</li>
                    </ul>

                    <p class="mt-4"><strong>To Marinate the Chicken:</strong></p>
                    <ul class="list-disc list-inside ml-4 space-y-1">
                        <li>340 g chicken breast (cut into 3/4-inch/2cm cubes)</li>
                        <li>1 teaspoon vegetable oil</li>
                        <li>1 teaspoon cornstarch</li>
                        <li>1 teaspoon shaoxing wine</li>
                        <li>1/8 teaspoon salt</li>
                        <li>1 pinch white pepper</li>
                    </ul>

                    <p class="mt-4"><strong>Sauce:</strong></p>
                    <ul class="list-disc list-inside ml-4 space-y-1">
                        <li>1 tablespoon light soy sauce</li>
                        <li>1/2 teaspoon dark soy sauce</li>
                        <li>1 tablespoon rice wine vinegar</li>
                        <li>1 teaspoon sugar</li>
                        <li>45 ml water</li>
                        <li>1 teaspoon cornstarch</li>
                    </ul>

                    <p class="mt-4"><strong>Stir Fry:</strong></p>
                    <ul class="list-disc list-inside ml-4 space-y-1">
                        <li>45 ml vegetable oil</li>
                        <li>3 cloves garlic (smashed and sliced)</li>
                        <li>2 thin slices ginger (minced)</li>
                        <li>2 dried red chillies (deseeded and chopped; these can be quite spicy, so adjust according to
                            your own tastes)</li>
                        <li>1/2 teaspoon sichuan peppercorn powder</li>
                        <li>6 scallions (white portions only, cut into 3/4-inch/2cm pieces)</li>
                    </ul>
                </div>

                <button
                    class="mt-6 px-4 py-2 border border-gray-300 rounded-full text-gray-700 hover:bg-gray-100 transition-colors flex items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round"
                            d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                    ADD TO SHOPPING LIST
                </button>
            </div>

            <!-- Instructions -->
            <div class="mb-8">
                <h2 class="text-xl font-bold text-gray-800 mb-4">Instructions</h2>

                <!-- Instruction Steps (Placeholder) -->
                <div class="space-y-4 text-gray-700 leading-relaxed">
                    <div>
                        <p class="font-semibold">Step 1</p>
                        <p>Input Text</p>
                    </div>
                    <div>
                        <p class="font-semibold">Step 2</p>
                        <p>Input Text</p>
                    </div>
                    <!-- Repeat for more steps or use v-for -->
                </div>
            </div>


        </div>
    </div>
</template>

<script setup lang="ts">
// Script logic will go here
</script>

<style scoped>
/* Component-specific styles */
</style>