<!-- components/Navbar.vue -->
<template>
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <!-- Left side - Brand logo -->
          <div class="flex items-center">
            <NuxtLink 
              to="/" 
              class="text-xl font-bold text-gray-900 hover:text-primary-600 transition-colors"
            >
              Dishcovery
            </NuxtLink>
          </div>
  
          <!-- Center - Search bar -->
          <!-- <div class="flex flex-1 items-center justify-center px-2 lg:ml-6 lg:justify-end">
            <div class="w-full max-w-lg lg:max-w-xs">
              <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <svg class="h-5 w-5 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
                  </svg>
                </div>
                <input
                  type="text"
                  placeholder="Search"
                  class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500 sm:text-sm"
                />
              </div>
            </div>
          </div> -->
  
          <!-- Right side - Navigation links -->
          <div class="flex items-center space-x-4">
            <NuxtLink
              to="/login"
              class="px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:text-primary-600 hover:bg-gray-50 transition-colors"
              active-class="text-primary-600 bg-gray-100"
            >
              Log in
            </NuxtLink>
            <NuxtLink
              to="/signup"
              class="px-3 py-2 rounded-md text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 transition-colors"
            >
              Register
            </NuxtLink>
          </div>
        </div>
      </div>
    </nav>
  </template>
  
