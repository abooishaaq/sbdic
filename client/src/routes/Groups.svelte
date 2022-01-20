<script>
    import Loading from "../components/Loading.svelte";
    
    let groups = [];
    let loading = true;

    fetch("/api/list-groups")
        .then((res) => {
            console.log(res);
            return res.json();
        })
        .then((res) => {
            if (res) {
                groups = res;
            }
            loading = false;
        });
</script>


<svelte:head>
    <title>Groups</title>
</svelte:head>
<div id="container">
    <h1>Groups</h1>
    {#if loading}
        <Loading />
    {:else if groups.length === 0}
        <p>No groups found. Create one <a href="/#/create-group">here</a></p>
    {:else}
        {#each groups as group}
            <a class="group" href={`/#/group/${encodeURIComponent(group.title)}`}>{group.title}</a>
        {/each}
    {/if}
</div>

<style>
    a.group {
        color: black;
        display: block;
        font-size: 1.2em;
        padding: 12px 0;
    }
</style>
