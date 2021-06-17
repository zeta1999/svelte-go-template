import Home from './routes/Home.svelte';
import Auth from './routes/Auth.svelte';
import Nf from './routes/Nf.svelte';

export default {
    '/': Home,
    '/auth/:jwt': Auth,
    // The catch-all route must always be last
    '*': Nf
};
