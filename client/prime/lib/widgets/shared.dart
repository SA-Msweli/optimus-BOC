import 'package:flutter/material.dart';
import '../theme.dart';

/// A reusable card container styled to the Optimus brand.
class InfoCard extends StatelessWidget {
  final String? title;
  final Widget child;
  final EdgeInsets padding;

  const InfoCard({
    super.key,
    this.title,
    required this.child,
    this.padding = const EdgeInsets.all(16),
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      margin: const EdgeInsets.only(bottom: 12),
      decoration: AppTheme.cardDecoration,
      child: Padding(
        padding: padding,
        child: title != null
            ? Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(title!, style: AppTheme.subheading),
                  const SizedBox(height: 10),
                  child,
                ],
              )
            : child,
      ),
    );
  }
}

/// A key-value row for displaying labelled data.
class KVRow extends StatelessWidget {
  final String label;
  final String value;
  final bool mono;

  const KVRow({
    super.key,
    required this.label,
    required this.value,
    this.mono = false,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 3),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: 140,
            child: Text(label, style: AppTheme.caption),
          ),
          Expanded(
            child: Text(
              value,
              style: mono ? AppTheme.monoValue : AppTheme.body,
            ),
          ),
        ],
      ),
    );
  }
}

/// Shows a red banner for error messages with dismiss.
class ErrorBanner extends StatelessWidget {
  final String message;
  final VoidCallback? onDismiss;

  const ErrorBanner({super.key, required this.message, this.onDismiss});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 10),
      margin: const EdgeInsets.only(bottom: 12),
      decoration: BoxDecoration(
        color: AppTheme.error.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: AppTheme.error.withValues(alpha: 0.3)),
      ),
      child: Row(
        children: [
          const Icon(Icons.error_outline, color: AppTheme.error, size: 20),
          const SizedBox(width: 8),
          Expanded(
              child: Text(message,
                  style: AppTheme.body.copyWith(color: AppTheme.error))),
          if (onDismiss != null)
            GestureDetector(
              onTap: onDismiss,
              child:
                  const Icon(Icons.close, color: AppTheme.error, size: 18),
            ),
        ],
      ),
    );
  }
}

/// Shows a green banner for success / tx hash feedback.
class SuccessBanner extends StatelessWidget {
  final String message;

  const SuccessBanner({super.key, required this.message});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.symmetric(horizontal: 14, vertical: 10),
      margin: const EdgeInsets.only(bottom: 12),
      decoration: BoxDecoration(
        color: AppTheme.secondary.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: AppTheme.secondary.withValues(alpha: 0.3)),
      ),
      child: Row(
        children: [
          const Icon(Icons.check_circle_outline,
              color: AppTheme.secondary, size: 20),
          const SizedBox(width: 8),
          Expanded(
              child: Text(message,
                  style: AppTheme.body.copyWith(color: AppTheme.secondary))),
        ],
      ),
    );
  }
}

/// A labelled loading indicator.
class LoadingOverlay extends StatelessWidget {
  final String? label;

  const LoadingOverlay({super.key, this.label});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          const CircularProgressIndicator(),
          if (label != null) ...[
            const SizedBox(height: 12),
            Text(label!, style: AppTheme.caption),
          ],
        ],
      ),
    );
  }
}

/// Truncates a hex address to 0x1234…abcd format.
String truncateAddress(String addr) {
  if (addr.length <= 12) return addr;
  return '${addr.substring(0, 6)}…${addr.substring(addr.length - 4)}';
}
