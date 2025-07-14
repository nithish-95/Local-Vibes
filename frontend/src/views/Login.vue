<template>
  <div class="flex items-center justify-center min-h-screen bg-light-bg px-4 py-8">
    <div class="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-xl">
      <button @click="$router.go(-1)" class="mb-4 text-neutral-600 hover:text-neutral-800 flex items-center font-medium transition-colors duration-300 font-sans">
        <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
        Back
      </button>
      <h2 class="text-3xl font-serif text-center text-light-text">Login to Your Account</h2>
      <Form @submit="login" :validation-schema="schema" :validate-on-mount="false" class="space-y-6">
        <div class="relative h-11 w-full min-w-[200px]">
          <Field type="text" name="username" id="username" autocomplete="username" 
            class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
            placeholder=" " />
          <label for="username"
            class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
            Username
          </label>
        </div>
        <ErrorMessage name="username" class="text-red-500 text-sm" />

        <div class="relative h-11 w-full min-w-[200px]">
          <Field type="password" name="password" id="password" autocomplete="current-password" 
            class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
            placeholder=" " />
          <label for="password"
            class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
            Password
          </label>
        </div>
        <ErrorMessage name="password" class="text-red-500 text-sm" />

        <div>
          <button type="submit" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-primary-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none w-full">Login</button>
        </div>
      </Form>
      <p class="text-sm text-center text-neutral-600 font-sans">
        Don't have an account? 
        <router-link to="/register" class="font-medium text-primary-500 hover:underline">Register</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import { useSessionStore } from '../stores/session';
import { useToast } from 'vue-toastification';
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';

export default {
  name: 'Login',
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    return {
      schema: yup.object({
        username: yup.string().required('Username is required'),
        password: yup.string().required('Password is required'),
      }),
    };
  },
  methods: {
    async login(values) {
      try {
        const sessionStore = useSessionStore();
        await sessionStore.login(values.username, values.password);
        this.$router.push('/');
      } catch (error) {
        console.error('Login failed:', error);
        const toast = useToast();
        toast.error('Login failed. Please check your username and password.');
      }
    },
  },
};
</script>