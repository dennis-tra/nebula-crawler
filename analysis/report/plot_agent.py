import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd

import lib_plot
from lib_db import DBClient, NodeClassification
from lib_fmt import fmt_barplot, fmt_thousands
from lib_agent import agent_name, go_ipfs_version, go_ipfs_v08_version


def main(client: DBClient):
    sns.set_theme()

    def plot_agent(results, plot_name):
        results_df = pd.DataFrame(results, columns=['agent_version', 'count']).assign(
            agent_name=lambda df: df.agent_version.apply(agent_name),
            go_ipfs_version=lambda df: df.agent_version.apply(go_ipfs_version),
            go_ipfs_v08_version=lambda df: df.agent_version.apply(go_ipfs_v08_version)
        )

        agent_names_df = results_df \
            .groupby(by='agent_name', as_index=False) \
            .sum() \
            .sort_values('count', ascending=False)
        agent_names_total = agent_names_df['count'].sum()

        go_ipfs_versions_df = results_df \
            .groupby(by='go_ipfs_version', as_index=False) \
            .sum() \
            .sort_values('count', ascending=False)
        go_ipfs_versions_total = go_ipfs_versions_df['count'].sum()

        go_ipfs_v08_versions_df = results_df \
            .groupby(by='go_ipfs_v08_version', as_index=False) \
            .sum() \
            .sort_values('count', ascending=False)
        go_ipfs_v08_versions_total = go_ipfs_v08_versions_df['count'].sum()

        # Plotting

        fig, (ax11, ax21, ax31) = plt.subplots(1, 3, figsize=(15, 5))  # rows, cols

        sns.barplot(ax=ax11, x='agent_name', y='count', data=agent_names_df)
        fmt_barplot(ax11, agent_names_df["count"], agent_names_total)
        ax11.title.set_text(f"Agent Types (Total {fmt_thousands(agent_names_total)})")
        ax11.set_xlabel("Agent")
        ax11.set_ylabel("Count")

        sns.barplot(ax=ax21, x='go_ipfs_version', y='count', data=go_ipfs_versions_df)
        fmt_barplot(ax21, go_ipfs_versions_df["count"], go_ipfs_versions_total)
        ax21.title.set_text(f"go-ipfs Versions Distribution (Total {fmt_thousands(go_ipfs_versions_total)})")
        ax21.set_xlabel("go-ipfs Version")
        ax21.set_ylabel("Count")

        sns.barplot(ax=ax31, x='go_ipfs_v08_version', y='count', data=go_ipfs_v08_versions_df)
        fmt_barplot(ax31, go_ipfs_v08_versions_df["count"], go_ipfs_v08_versions_total)
        ax31.title.set_text(f"go-ipfs v0.8.0 Versions Distribution (Total {fmt_thousands(go_ipfs_v08_versions_total)})")
        ax31.set_xlabel("go-ipfs v0.8.0 Version")
        ax31.set_ylabel("Count")

        plt.suptitle(f"Agent Version Distribution of '{plot_name}' Peers")
        plt.tight_layout()
        lib_plot.savefig(f"agents-{plot_name}")
        plt.show()

    results = client.get_agent_versions_distribution()
    plot_agent(results, "all")

    node_classes = [
        NodeClassification.DANGLING,
        NodeClassification.ONLINE,
        NodeClassification.ONEOFF,
        NodeClassification.ENTERED
    ]
    for node_class in node_classes:
        peer_ids = client.node_classification_funcs[node_class]()
        results = client.get_agent_versions_for_peer_ids(peer_ids)
        plot_agent(results, node_class.value)


if __name__ == '__main__':
    client = DBClient()
    main(client)
