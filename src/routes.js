import showJobsList from './components/ShowJobslist.vue';
import showJob from './components/ShowJob.vue';
import AddJob from './components/AddJob.vue';
import ManageJob from './components/Managejob.vue'

export default [
    { path: '/', component: showJobsList },
    { path: '/job/:id', component: showJob },
    { path: '/manage', component: AddJob },
    { path: '/manage-job', component: ManageJob }
]