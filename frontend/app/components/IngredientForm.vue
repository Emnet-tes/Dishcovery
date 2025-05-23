<template>
     <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">Ingredients <span class="text-red-500">*</span></label>
        <p class="text-sm text-gray-600 mb-4">List one ingredient per line, specifying quantities (1, 2), measurements (cups, spoons), and any prep details (chopped, sifted) along with the item. Let your creativity flow in every detail!</p>

        <div class="space-y-4">
          <!-- Example Ingredient Row -->
          <div class="flex items-center gap-4">
            <input type="text" placeholder="Qty" class="w-16 px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none
             focus:ring-green-600 focus:border-transparent">
            <select class="px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Measurement</option>
                <!-- Add measurement options here -->
            </select>
            <input type="text" placeholder="Item" class="flex-grow px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
            <button class="text-gray-500 hover:text-red-600 transition-colors">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          <!-- Repeat similar structure for more ingredients or use v-for -->
        </div>

        <button class="mt-4 text-green-700 flex items-center hover:text-green-800 transition-colors">
          <Icon name="uil:plus" class="h-5 w-5 mr-2" />
           Add ingredient
        </button>
      </div>

    </template>