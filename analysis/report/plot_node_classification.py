import seaborn as sns
import matplotlib.pyplot as plt

import lib_plot
from lib_db import DBClient, NodeClassification
from lib_fmt import fmt_thousands, thousands_ticker_formatter, fmt_percentage
from collections import OrderedDict


def main(db_client: DBClient):
    sns.set_theme()

    all_peer_ids = db_client.get_all_peer_ids()

    data = OrderedDict()
    for node_class in NodeClassification:
        get_peer_ids = db_client.node_classification_funcs[node_class]
        data[node_class.value] = len(get_peer_ids())

    data = OrderedDict(reversed(sorted(data.items(), key=lambda item: item[1])))  # order dict by count decreasing

    fig, (ax) = plt.subplots(figsize=(10, 5))
    sns.barplot(ax=ax, x=list(data.keys()), y=list(data.values()))
    ax.get_yaxis().set_major_formatter(thousands_ticker_formatter)
    ax.bar_label(ax.containers[0], list(map(fmt_percentage(len(all_peer_ids)), data.values())))
    ax.title.set_text(f"Node Classification of {fmt_thousands(len(all_peer_ids))} Visited Peers")
    ax.set_ylabel("Count")

    plt.tight_layout()
    lib_plot.savefig("node_classifications")
    plt.show()


if __name__ == '__main__':
    db_client = DBClient()
    main(db_client)
