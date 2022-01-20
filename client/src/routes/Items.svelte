<script>
    import ItemList from "../components/ItemList.svelte";
</script>

<svelte:head>
    <title>Item List</title>
</svelte:head>

<div id="container">
    <ItemList title="Items" url="/api/list-items" />
</div>
