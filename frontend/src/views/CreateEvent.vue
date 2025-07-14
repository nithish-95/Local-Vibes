<template>
  <div class="max-w-2xl mx-auto p-8 bg-white rounded-lg shadow-xl">
    <button @click="$router.go(-1)" class="mb-4 text-neutral-600 hover:text-neutral-800 flex items-center font-medium transition-colors duration-300 font-sans">
      <svg class="w-5 h-5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
      Back
    </button>
    <h2 class="text-3xl font-serif text-center text-light-text mb-8">Create Event</h2>
    <Form @submit="create" :validation-schema="schema" class="space-y-6">
      <div class="relative h-11 w-full min-w-[200px]">
        <Field type="text" name="title" id="title"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" " />
        <label for="title"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Title
        </label>
        <ErrorMessage name="title" class="text-red-500 text-sm" />
      </div>
      <div class="relative w-full min-w-[200px]">
        <Field as="textarea" name="description" id="description" rows="4"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" "></Field>
        <label for="description"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Description
        </label>
        <ErrorMessage name="description" class="text-red-500 text-sm" />
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="relative h-11 w-full min-w-[200px]">
          <Field type="date" name="date" id="date"
            class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50" />
          <label for="date"
            class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
            Date
          </label>
          <ErrorMessage name="date" class="text-red-500 text-sm" />
        </div>
        <div class="relative h-11 w-full min-w-[200px]">
          <Field type="time" name="time" id="time"
            class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50" />
          <label for="time"
            class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
            Time
          </label>
          <ErrorMessage name="time" class="text-red-500 text-sm" />
        </div>
      </div>
      <div class="relative h-11 w-full min-w-[200px]">
        <Field type="text" name="location" id="location"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" " />
        <label for="location"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Location
        </label>
        <ErrorMessage name="location" class="text-red-500 text-sm" />
      </div>
      <div class="relative h-11 w-full min-w-[200px]">
        <Field type="number" name="capacity" id="capacity"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" " />
        <label for="capacity"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Capacity
        </label>
        <ErrorMessage name="capacity" class="text-red-500 text-sm" />
      </div>
      <div class="relative h-11 w-full min-w-[200px]">
        <Field type="url" name="image_url" id="image_url" autocomplete="off"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" " />
        <label for="image_url"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Event Image URL (Optional)
        </label>
        <ErrorMessage name="image_url" class="text-red-500 text-sm" />
      </div>
      <div class="relative w-full min-w-[200px]">
        <Field as="textarea" name="rulesInput" id="rules" autocomplete="off" rows="4"
          class="peer h-full w-full rounded-md border border-neutral-200 border-t-transparent bg-transparent px-3 py-3 font-sans text-sm font-normal text-neutral-700 outline outline-0 transition-all placeholder-shown:border placeholder-shown:border-neutral-200 placeholder-shown:border-t-neutral-200 focus:border-2 focus:border-accent-500 focus:border-t-transparent focus:outline-0 disabled:border-0 disabled:bg-neutral-50"
          placeholder=" "></Field>
        <label for="rules"
          class="before:content[' '] after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-neutral-400 transition-all before:pointer-events-none before:mt-[6.5px] before:mr-1 before:box-border before:block before:h-1.5 before:w-2.5 before:rounded-tl-md before:border-t before:border-l before:border-neutral-200 before:transition-all after:pointer-events-none after:mt-[6.5px] after:ml-1 after:box-border after:block after:h-1.5 after:w-2.5 after:rounded-tr-md after:border-t after:border-r after:border-neutral-200 after:transition-all peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.1] peer-placeholder-shown:text-neutral-500 peer-placeholder-shown:before:border-transparent peer-placeholder-shown:after:border-transparent peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-accent-500 peer-focus:before:border-t-2 peer-focus:before:border-l-2 peer-focus:before:border-accent-500 peer-focus:after:border-t-2 peer-focus:after:border-r-2 peer-focus:after:border-accent-500 peer-disabled:text-transparent peer-disabled:before:border-transparent peer-disabled:after:border-transparent peer-disabled:peer-placeholder-shown:text-neutral-500">
          Custom Rules (one per line)
        </label>
        <ErrorMessage name="rulesInput" class="text-red-500 text-sm" />
      </div>
      <div>
        <button type="submit" class="align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg bg-primary-500 text-dark-text shadow-md hover:shadow-lg focus:opacity-[0.85] focus:shadow-none active:opacity-[0.85] active:shadow-none w-full">Create Event</button>
      </div>
    </Form>
  </div>
</template>

<script>
import { createEvent } from '../api/events';
import { useToast } from 'vue-toastification';
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';

export default {
  name: 'CreateEvent',
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  setup() {
    const toast = useToast();
    return { toast };
  },
  data() {
    return {
      schema: yup.object({
        title: yup.string().required('Title is required').max(25, 'Title cannot exceed 25 characters'),
        description: yup.string().required('Description is required').max(500, 'Description cannot exceed 500 characters'),
        date: yup.date().required('Date is required').nullable().min(new Date(), 'Event date cannot be in the past'),
        time: yup.string().required('Time is required').test(
          'is-future-time',
          'Event must be at least 1 day from now',
          function (value) {
            const { date } = this.parent;
            if (!date || !value) return true; // Let other validations handle missing date/time

            const eventDateTime = new Date(date);
            const [hours, minutes] = value.split(':').map(Number);
            eventDateTime.setHours(hours, minutes, 0, 0);

            const now = new Date();
            const oneDayFromNow = new Date(now.getTime() + 24 * 60 * 60 * 1000); // Add 1 day in milliseconds

            return eventDateTime > oneDayFromNow;
          }
        ),
        location: yup.string().required('Location is required'),
        capacity: yup.number().required('Capacity is required').min(1, 'Capacity must be at least 1'),
        image_url: yup.string().url('Must be a valid URL').nullable(),
        rulesInput: yup.string().nullable(),
      }),
      rulesInput: '',
    };
  },
  methods: {
    async create(values) {
      try {
        const eventData = { ...values };
        eventData.rules = this.rulesInput.split('\n').map(rule => rule.trim()).filter(rule => rule.length > 0);
        await createEvent(eventData);
        this.toast.success('Event created successfully!');
        this.$router.push('/');
      } catch (error) {
        console.error(error);
        this.toast.error('Error creating event: ' + (error.response?.data?.error || error.message));
      }
    },
  },
};
</script>
