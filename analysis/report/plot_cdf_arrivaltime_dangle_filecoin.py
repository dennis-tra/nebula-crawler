import pandas as pd
import numpy as np
import seaborn as sns
from matplotlib import pyplot as plt, ticker

from analysis.report import lib_plot
from analysis.report.lib_db_filecoin import DBClientFilecoin
from analysis.report.lib_fmt import fmt_thousands
from lib_db import DBClient


def main(client: DBClient):
    sns.set_theme()

    peer_ids = client.get_dangling_peer_ids()
    arrivals = client.get_inter_arrival_time(peer_ids)

    results_df = pd.DataFrame(arrivals, columns=['id', 'peer_id', 'diff_in_s'])
    results_df = results_df.assign(
        diff_in_h=results_df.diff_in_s.apply(lambda x: x / 3600),
    )

    fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(15, 5), sharey=True)

    sns.ecdfplot(ax=ax1, x="diff_in_h", data=results_df)

    ax1.set_xlim(0, 96)
    ax1.set_xticks(np.arange(0, 100, step=4))
    ax1.set_xlabel("Time in Hours")
    ax1.set_ylabel("Number of Peers in %")
    ax1.get_yaxis().set_major_formatter(ticker.FuncFormatter(lambda x, p: "%d" % int(x * 100)))
    ax1.legend(loc='lower right', labels=[f"dangling ({fmt_thousands(len(results_df))})"])

    ax1.title.set_text(f"CDF of Inter Arrival Times of Dangling Peers")

    results = client.get_agent_versions_distribution()
    top_agent_versions = [result[0] for result in results[:4]]

    labels = []
    for agent in top_agent_versions:
        peer_ids = client.get_peer_ids_for_agent_versions([agent])
        arrivals = client.get_inter_arrival_time(peer_ids)
        data = pd.DataFrame(arrivals, columns=['id', 'peer_id', 'diff_in_s'])
        data = data.assign(
            diff_in_h=data.diff_in_s.apply(lambda x: x / 3600),
        )
        labels += [f"{agent} ({fmt_thousands(len(data))})"]
        sns.ecdfplot(ax=ax2, x="diff_in_h", data=data)
        ax2.set_xlim(0, 96)
        ax2.set_xticks(np.arange(0, 100, step=4))
        ax2.set_xlabel("Time in Hours")
        ax2.set_ylabel("Number of Peers in %")
        ax2.get_yaxis().set_major_formatter(ticker.FuncFormatter(lambda x, p: "%d" % int(x * 100)))

    ax2.title.set_text(f"CDF of Inter Arrival Times by Agent")
    ax2.legend(loc='lower right', labels=labels)

    plt.tight_layout()
    lib_plot.savefig("cdf-inter-arrival-dangling-filecoin")
    plt.show()


if __name__ == '__main__':
    client = DBClientFilecoin()
    main(client)
