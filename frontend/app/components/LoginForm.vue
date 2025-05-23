<script setup lang="ts">
import { ref } from 'vue';
import { useForm, useField } from 'vee-validate';
import { toTypedSchema } from "@vee-validate/zod"
import { object, string} from 'zod';

const validationSchema = toTypedSchema(
object({
  email: string()
    .min(1, { message: "Email is required" })
    .email({ message: "Invalid email address" }),
  password: string()
.min(8, { message: "Must be at least 8 characters" })
.regex(/[a-z]/, { message: "Must include at least one lowercase letter" })
.regex(/[A-Z]/, { message: "Must include at least one uppercase letter" })
.regex(/[0-9]/, { message: "Must include at least one number" })
.regex(/[^A-Za-z0-9]/, { message: "Must include at least one special character" })
})
);

// Initialize the form with the validation schema
const { handleSubmit ,errors} = useForm({
  validationSchema,
});

const { value: email } = useField("email");
const { value: password } = useField("password");
const router = useRouter();
const passwordFieldType = ref('password'); 

const togglePasswordVisibility = () => {
  passwordFieldType.value = passwordFieldType.value === 'password' ? 'text' : 'password';
};

const onSubmit = handleSubmit((values) => {

  router.push("/");


});
</script>
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

        <div class="bg-white p-8 rounded-lg shadow-sm ">
          <h2 class="text-xl font-semibold text-gray-900 mb-6">
            Log In
          </h2>

          <!-- Form Fields -->
          <form class="space-y-5 rounded-l-md">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Email</label
              >
              <input
                type="email"
                v-model="email"
                id="email"
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
                v-model="password"
                id="password"
                placeholder="Enter password"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-green-500"
              />
            </div>

            <button
              type="submit"
              class="w-full py-2 px-4 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2"
            >
              Log In
            </button>
          </form>

          <div class="mt-4 text-center text-sm text-gray-600">
            Don't have an account?
            <NuxtLink 
              to="/auth/register" 
             class="text-green-600 hover:text-green-500 font-medium"
              >Register</NuxtLink
            >
          </div>
        </div>
      </div>
    </div>
</template>