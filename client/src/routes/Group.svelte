<script>
    import ItemList from "../components/ItemList.svelte";
    import Loading from "../components/Loading.svelte";

    const splittedHash = window.location.hash.split("/");
    const encodedGroup = splittedHash[splittedHash.length - 1];
    let group = decodeURIComponent(encodedGroup);

    let loading = true;
    let id = -1;
    let desc = "";

    fetch(`/api/get-group?title=${encodedGroup}`)
        .then((res) => res.json())
        .then((res) => {
            if (res.err) alert(res.err);
            else {
                id = res.id;
                desc = res.desc;
            }
            loading = false;
        })
        .catch(console.log);

    const deleteGroup = () => {
        if(confirm("Are you sure ou want to delete this group? This won't delete its items.")) {
            fetch(`/api/delete-group?id=${id}`)
            .then((res) => res.json())
            .then((res) => {
                if (res.err) alert(res.err);
                location.hash = "#";
            })
            .catch(console.log);
        }
    };
</script>

<svelte:head>
    <title>Group {group}</title>
</svelte:head>
<div id="container">
    {#if loading}
        <Loading />
    {:else if id == -1}
        <p>Group does not exist.</p>
    {:else}
        <h1>{group}</h1>
        <p>{desc}</p>
        <ItemList title="" url={`/api/get-group-items?id=${id}`} />
        <button on:click={deleteGroup}>Delete group</button>
    {/if}
</div>

<style>
    button {
        color: #eee;
        max-width: 320px;
        background-color: crimson;
        margin-top: 24px;
    }
</style>
