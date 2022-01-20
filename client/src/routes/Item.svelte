<script>
    import Item from "../components/Item.svelte";
    import Loading from "../components/Loading.svelte";

    const splittedHash = window.location.hash.split("/");
    const encodedTitle = splittedHash[splittedHash.length - 1];
    let title = decodeURIComponent(encodedTitle);

    let desc = "",
        count = 0,
        tags = [],
        id = "",
        created_at = "",
        updated_at = "";

    let loading = true;

    $: fetch(`/api/get-item?title=${title}`)
        .then((res) => {
            console.log(res);
            if (res.status == 400) {
                title = "";
                loading = false
                return;
            }
            return res.json();
        })
        .then((res) => {
            if (res) {
                if (typeof res.tags === "undefined") res.tags = [];
                ({ desc, count, tags, id, created_at, updated_at } = res);
                loading = false;
            }
        });
</script>


<svelte:head>
    <title>{title}</title>
</svelte:head>
<div id="container">
    {#if loading}
        <Loading />
    {:else if title.length > 0}
        <Item {title} {desc} {count} {tags} {id} {created_at} {updated_at} />
    {:else}
        <p>Item does not exist.</p>
    {/if}
</div>

<style>
        
</style>
