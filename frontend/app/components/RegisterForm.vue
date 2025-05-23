<template>
    <div
      class="w-full lg:w-1/2 bg-gray-50 flex items-center justify-center p-8 md:rounded-l-3xl "
    >
      <div class="w-full max-w-md">
        <div class="lg:hidden text-center mb-8">
          <h1 class="text-3xl font-bold text-gray-900 mb-2">Kitchen Tales</h1>
          <p class="text-gray-600">
            Embark on a culinary journey with us!<br />
            Sign up to unlock a world of delicious recipes, and personalized
            cooking experiences.
          </p>
        </div>

        <div class="bg-white p-8 rounded-lg shadow-sm">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">
            Create an Account
          </h2>
          <!-- Form Fields -->
          <form class="space-y-5 rounded-l-md">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Full Name</label
              >
              <input
                type="text"
                placeholder="Enter your full name"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Email</label
              >
              <input
                type="email"
                placeholder="Enter your e-mail"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Password</label
              >
              <input
                type="password"
                placeholder="Enter password"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <button
              type="submit"
              class="w-full py-2 px-4 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            >
              Create Account
            </button>
          </form>

          <div class="mt-4 text-center text-sm text-gray-600">
            Already have an account?
            <NuxtLink 
              to="/auth/login" 
             class="text-green-600 hover:text-green-500 font-medium"
              >Log in</NuxtLink>
          </div>
        </div>
      </div>
    </div>
</template>