import Home from './routes/Home.svelte';
import Lorem from './routes/Lorem.svelte';
import Nf from './routes/Nf.svelte';

export default {
    '/': Home,
    '/lorem/:repeat': Lorem,
    // The catch-all route must always be last
    '*': Nf
};
