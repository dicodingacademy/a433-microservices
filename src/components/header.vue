<template>
    <nav class="mt-5 mb-20">
        <ul class="list-none flex justify-between">
            <li>
                <router-link class="mr-4 flex items-center justify-between" to="/" exact>
                <svg class="w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                </svg>
                    Home
                </router-link>
            </li>
            <li>
                <router-link class="flex items-center justify-between" to="/manage" exact>
                <svg class="w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                    Manage jobs
                </router-link>
            </li>
        </ul>
    </nav>
</template>

<script>
export default {}
</script>

<style scoped>
a {
    --tw-text-opacity: 1;
    color: rgba(248, 113, 113, var(--tw-text-opacity));
}
.router-link-active{
    --tw-text-opacity: 1;
    color: rgba(55, 65, 81, var(--tw-text-opacity));
}
</style>