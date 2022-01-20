<script>
    import Home from "./routes/Home.svelte";
    import Item from "./routes/Item.svelte";
    import Items from "./routes/Items.svelte";
    import Group from "./routes/Group.svelte";
    import Groups from "./routes/Groups.svelte";
    import Tag from "./routes/Tag.svelte";
    import AddItem from "./routes/AddItem.svelte";
    import CreateGroup from "./routes/CreateGroup.svelte";
    
    
    let hash = window.location.hash;

    setInterval(() => {
        hash = window.location.hash;
    }, 200);

    let Component = Home;

    $: console.log(hash);
    $: switch (hash.match(/#\/(\w|-)*/)?.[0]) {
        case "#/":
            Component = Home;
            break;
        case "#/tag":
            Component = Tag;
            break;
        case "#/item":
            Component = Item;
            break;
        case "#/group":
            Component = Group;
            break;
        case "#/items":
            Component = Items;
            break;
        case "#/groups":
            Component = Groups;
            break;
        case "#/add-item":
            Component = AddItem;
            break;
        case "#/create-group":
            Component = CreateGroup;
            break;
        default:
            Component = Home;
    }
</script>

<nav>
    <a href="/#/">Home</a>
    <a href="/#/items">Items</a>
    <a href="/#/groups">Groups</a>
    <a href="/#/add-item">Add&nbsp;Item</a>
    <a href="/#/create-group">Create&nbsp;Group</a>
</nav>
<div>
    <svelte:component this={Component} />
</div>

<style>
    div {
        flex-grow: 3;
        display: flex;
        justify-content: center;
        align-items: flex-start;
        padding: 10vh 12px;
        min-height: 100vh;
    }
    nav {
        padding: 12px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: right;
        height: 100%;
        font-size: 1.1em;
    }
</style>
