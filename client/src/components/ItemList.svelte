<script>
    import Loading from "../components/Loading.svelte";
    import ItemMin from "./ItemMin.svelte";

    export let url = "";
    export let title = "";
    export let fallback = "#";

    let items = [];
    let loading = true;

    fetch(url)
        .then((res) => {
            console.log(res);
            if (res.status == 400) {
                title = "";
                loading = false;
                return;
            }
            return res.json();
        })
        .then((res) => {
            if (res) {
                if(res.err)  {
                    alert(res.err);
                    location.hash = fallback;
                }
                else items = res;
                loading = false;
            }
        });
</script>

<div>
    {#if title && title.length > 0}
        <h1>{title}</h1>
    {/if}
    {#if loading}
        <Loading />
    {:else}
        {#each items as item}
            <ItemMin {item} />
        {/each}
    {/if}
</div>
