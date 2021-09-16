from multiaddr import Multiaddr
import helper
from icmplib import async_multiping
import asyncio


# check_node_ping_async gets the if icmp ping works to of given nodes.
# It takes an sql connection, the peer ids as the arguments, and
# returns a map indicating if icmp ping works for the given peers.
async def check_node_ping_async(conn, peer_ids):
    cur = conn.cursor()
    res = dict()
    cur.execute(
        """
        SELECT id, multi_addresses
        FROM peers
        WHERE id IN (%s)
        """ % ','.join(['%s'] * len(peer_ids)),
        tuple(peer_ids)
    )
    ids = []
    temp = []
    for id, maddr_strs in cur.fetchall():
        addrs = []
        for maddr_str in maddr_strs:
            maddr = Multiaddr(maddr_str)
            try:
                address = helper.node_address(maddr)
                addrs.append(address)
            except:
                pass
        if len(addrs) == 0:
            res[id] = False
        else:
            ids.append(id)
            temp.append(async_multiping(addrs))
    hosts = await asyncio.gather(*temp)
    for i in range(len(hosts)):
        found = False
        for h in hosts[i]:
            if h.is_alive:
                found = True
                break
        if found:
            res[ids[i]] = True
        else:
            res[ids[i]] = False
    return res


# check_node_ping gets the if icmp ping works to of given nodes.
# It takes an sql connection, the peer ids as the arguments, and
# returns a map indicating if icmp ping works for the given peers.
def check_node_ping(conn, peer_ids):
    return asyncio.run(check_node_ping_async(conn, peer_ids))
